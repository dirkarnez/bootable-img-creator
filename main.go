package main

import (
	"log"
	"github.com/diskfs/go-diskfs"
	"github.com/diskfs/go-diskfs/disk"
)

/*
    // 1. 在 Windows 上直接建立一個 500MB 的空白大檔案（取代 fallocate/dd）
    var diskSize int64 = 500 * 1024 * 1024
    d, _ := diskfs.Create("sdcard.img", diskSize, diskfs.SectorSize512)

    // 2. 直接在檔案內部寫入 MBR 分割表
    table := &mbr.Table{
        Partitions: []mbr.Partition{
            {Bootable: true, Type: mbr.Fat16, StartReader: 2048, Size: 65536}, // Boot 分區
            {Bootable: false, Type: mbr.Linux, StartReader: 67584, Size: 400000}, // Rootfs 分區
        },
    }
    d.Partition(table)

    // 3. 在指定分區內直接寫入 FAT 檔案系統，並把 Windows 本地的 zImage 塞進去
    fs, _ := d.CreateFilesystem(disk.FilesystemSpec{Partition: 1, FSType: diskfs.TypeFat32})
    f, _ := fs.OpenFile("/zImage", os.O_CREATE|os.O_WRONLY)
    // ...隨後直接利用 Go 的 io.Copy 將核心寫入，完全不需要 mount
*/
func main() {
	// 1. 打開你已經創建好的映像檔（或是新創建的檔案）
	var diskSize int64 = 1 * 1024 // 10 KB
	mydisk, err := diskfs.Create("output.img", diskSize, diskfs.SectorSize4k)
	if err != nil {
		panic(err)
	}

	ddWithSeek(mydisk, []byte("Hello, World!!!"), 24)

	// // 2. 打開你要寫入的 U-Boot 原始二進位檔
	// uBootFile, err := os.Open("u-boot-sunxi-with-spl.bin")
	// if err != nil {
	// 	panic(err)
	// }
	// defer uBootFile.Close()

	// br := bytes.NewReader()


}


func ddWithSeek(disk *disk.Disk, bytes []byte, offset int64) (n int, err error) {
	w, err := disk.Backend.Writable()
	if err != nil {
		log.Fatalf("backend writable: %v", err)
	}

	return w.WriteAt(bytes, offset)
}


	// fspec := disk.FilesystemSpec{Partition: 0, FSType: filesystem.TypeSquashfs, VolumeLabel: "label"}
	// fs, err := mydisk.CreateFilesystem(fspec)
	// check(err)
	// defer func() {
	// 	if err := fs.Close(); err != nil {
	// 		check(err)
	// 	}
	// }()
	// rw, err := fs.OpenFile("demo.txt", os.O_CREATE|os.O_RDWR)
	// check(err)
	// content := []byte("demo")
	// _, err = rw.Write(content)
	// check(err)
	// sqs, ok := fs.(*squashfs.FileSystem)
	// if !ok {
	// 	check(fmt.Errorf("not a squashfs filesystem"))
	// }
	// err = sqs.Finalize(squashfs.FinalizeOptions{})
	// check(err)