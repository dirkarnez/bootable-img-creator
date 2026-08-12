package main

import (
	"fmt"
	"io"
	"os"

	"github.com/diskfs/go-diskfs"
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
	diskImg, err := diskfs.Open("sdcard.img")
	if err != nil {
		panic(err)
	}
	defer diskImg.Close()

	// 2. 打開你要寫入的 U-Boot 原始二進位檔
	uBootFile, err := os.Open("u-boot-sunxi-with-spl.bin")
	if err != nil {
		panic(err)
	}
	defer uBootFile.Close()

	// 3. 計算你的偏移量（例如 dd seek=8 且 bs=1k，代表 8KB = 8192 位元組）
	var offset int64 = 8 * 1024 

	// 4. 關鍵步驟：調用 Seek 來移動檔案指標 (whence 0 代表從檔案開頭計算)
	_, err = diskImg.Seek(offset, 0)
	if err != nil {
		panic(err)
	}

	// 5. 寫入資料：將 U-Boot 內容直接複製進去（此時指標已在 8KB 處）
	written, err := io.Copy(diskImg, uBootFile)
	if err != nil {
		panic(err)
	}

	fmt.Printf("成功在偏移量 %d 處寫入 %d 位元組的 U-Boot 原始資料\n", offset, written)
}
