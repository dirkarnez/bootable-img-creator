package main

import (
	"fmt"
	"io"
	"os"

	"github.com/diskfs/go-diskfs"
)

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
