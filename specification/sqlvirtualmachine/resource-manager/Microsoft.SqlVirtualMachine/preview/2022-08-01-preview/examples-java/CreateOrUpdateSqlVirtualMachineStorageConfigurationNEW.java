
import com.azure.resourcemanager.sqlvirtualmachine.models.DiskConfigurationType;
import com.azure.resourcemanager.sqlvirtualmachine.models.SqlStorageSettings;
import com.azure.resourcemanager.sqlvirtualmachine.models.SqlTempDbSettings;
import com.azure.resourcemanager.sqlvirtualmachine.models.StorageConfigurationSettings;
import com.azure.resourcemanager.sqlvirtualmachine.models.StorageWorkloadType;
import java.util.Arrays;

/**
 * Samples for SqlVirtualMachines CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/preview/2022-08-01-preview/examples/
     * CreateOrUpdateSqlVirtualMachineStorageConfigurationNEW.json
     */
    /**
     * Sample code: Creates or updates a SQL virtual machine for Storage Configuration Settings to NEW Data, Log and
     * TempDB storage pool.
     * 
     * @param manager Entry point to SqlVirtualMachineManager.
     */
    public static void
        createsOrUpdatesASQLVirtualMachineForStorageConfigurationSettingsToNEWDataLogAndTempDBStoragePool(
            com.azure.resourcemanager.sqlvirtualmachine.SqlVirtualMachineManager manager) {
        manager.sqlVirtualMachines().define("testvm").withRegion("northeurope").withExistingResourceGroup("testrg")
            .withVirtualMachineResourceId(
                "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Compute/virtualMachines/testvm")
            .withStorageConfigurationSettings(new StorageConfigurationSettings()
                .withSqlDataSettings(
                    new SqlStorageSettings().withLuns(Arrays.asList(0)).withDefaultFilePath("F:\\folderpath\\"))
                .withSqlLogSettings(
                    new SqlStorageSettings().withLuns(Arrays.asList(1)).withDefaultFilePath("G:\\folderpath\\"))
                .withSqlTempDbSettings(new SqlTempDbSettings().withDataFileSize(256).withDataGrowth(512)
                    .withLogFileSize(256).withLogGrowth(512).withDataFileCount(8).withDefaultFilePath("D:\\TEMP"))
                .withSqlSystemDbOnDataDisk(true).withDiskConfigurationType(DiskConfigurationType.NEW)
                .withStorageWorkloadType(StorageWorkloadType.OLTP))
            .create();
    }
}
