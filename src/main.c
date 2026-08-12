#include <stdio.h>
#include <stdlib.h>

#define SECTOR_SIZE 512
#define TOTAL_SECTORS 2880 // 2880 * 512 = 1.44 MB (Standard Floppy Size)

int main() {
    FILE *img_file = fopen("blank_disk.img", "wb");
    if (img_file == NULL) {
        perror("Error creating image file");
        return EXIT_FAILURE;
    }

    // Create a zero-filled buffer representing one sector
    unsigned char sector_buffer[SECTOR_SIZE] = {0};

    printf("Creating blank disk image...\n");

    // Write the empty sectors sequentially to the file
    for (int i = 0; i < TOTAL_SECTORS; i++) {
        size_t written = fwrite(sector_buffer, 1, SECTOR_SIZE, img_file);
        if (written != SECTOR_SIZE) {
            perror("Error writing sector data");
            fclose(img_file);
            return EXIT_FAILURE;
        }
    }

    fclose(img_file);
    printf("Successfully created 'blank_disk.img' (%ld bytes).\n", (long)TOTAL_SECTORS * SECTOR_SIZE);
    return EXIT_SUCCESS;
}
