
import com.azure.resourcemanager.sqlvirtualmachine.models.DiskConfigurationType;
import com.azure.resourcemanager.sqlvirtualmachine.models.SqlStorageSettings;
import com.azure.resourcemanager.sqlvirtualmachine.models.StorageConfigurationSettings;
import java.util.Arrays;

/**
 * Samples for SqlVirtualMachines CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/preview/2022-08-01-preview/examples/
     * CreateOrUpdateSqlVirtualMachineStorageConfigurationEXTEND.json
     */
    /**
     * Sample code: Creates or updates a SQL virtual machine for Storage Configuration Settings to EXTEND Data, Log or
     * TempDB storage pool.
     * 
     * @param manager Entry point to SqlVirtualMachineManager.
     */
    public static void
        createsOrUpdatesASQLVirtualMachineForStorageConfigurationSettingsToEXTENDDataLogOrTempDBStoragePool(
            com.azure.resourcemanager.sqlvirtualmachine.SqlVirtualMachineManager manager) {
        manager.sqlVirtualMachines().define("testvm").withRegion("northeurope").withExistingResourceGroup("testrg")
            .withVirtualMachineResourceId(
                "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Compute/virtualMachines/testvm")
            .withStorageConfigurationSettings(new StorageConfigurationSettings()
                .withSqlDataSettings(new SqlStorageSettings().withLuns(Arrays.asList(2)))
                .withDiskConfigurationType(DiskConfigurationType.EXTEND))
            .create();
    }
}
