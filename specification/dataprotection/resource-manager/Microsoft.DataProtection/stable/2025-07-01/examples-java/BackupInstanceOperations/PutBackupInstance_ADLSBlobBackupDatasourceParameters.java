
import com.azure.resourcemanager.dataprotection.models.AdlsBlobBackupDatasourceParameters;
import com.azure.resourcemanager.dataprotection.models.BackupInstance;
import com.azure.resourcemanager.dataprotection.models.Datasource;
import com.azure.resourcemanager.dataprotection.models.DatasourceSet;
import com.azure.resourcemanager.dataprotection.models.PolicyInfo;
import com.azure.resourcemanager.dataprotection.models.PolicyParameters;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for BackupInstances CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2025-07-01/examples/
     * BackupInstanceOperations/PutBackupInstance_ADLSBlobBackupDatasourceParameters.json
     */
    /**
     * Sample code: Create BackupInstance With ADLSBlobBackupDatasourceParameters.
     * 
     * @param manager Entry point to DataProtectionManager.
     */
    public static void createBackupInstanceWithADLSBlobBackupDatasourceParameters(
        com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager.backupInstances().define("adlsstorageaccount-adlsstorageaccount-19a76f8a-c176-4f7d-819e-95157e2b0071")
            .withExistingBackupVault("adlsrg", "adlsvault").withTags(mapOf("key1", "fakeTokenPlaceholder"))
            .withProperties(new BackupInstance().withFriendlyName("adlsstorageaccount\\adlsbackupinstance")
                .withDataSourceInfo(new Datasource()
                    .withDatasourceType("Microsoft.Storage/storageAccounts/adlsBlobServices")
                    .withObjectType("Datasource")
                    .withResourceId(
                        "/subscriptions/54707983-993e-43de-8d94-074451394eda/resourceGroups/adlsrg/providers/Microsoft.Storage/storageAccounts/adlsstorageaccount")
                    .withResourceLocation("centraluseuap").withResourceName("adlsstorageaccount")
                    .withResourceType("microsoft.storage/storageAccounts").withResourceUri(
                        "/subscriptions/54707983-993e-43de-8d94-074451394eda/resourceGroups/adlsrg/providers/Microsoft.Storage/storageAccounts/adlsstorageaccount"))
                .withDataSourceSetInfo(new DatasourceSet()
                    .withDatasourceType("Microsoft.Storage/storageAccounts/adlsBlobServices")
                    .withObjectType("DatasourceSet")
                    .withResourceId(
                        "/subscriptions/54707983-993e-43de-8d94-074451394eda/resourceGroups/adlsrg/providers/Microsoft.Storage/storageAccounts/adlsstorageaccount")
                    .withResourceLocation("centraluseuap").withResourceName("adlsstorageaccount")
                    .withResourceType("microsoft.storage/storageAccounts").withResourceUri(
                        "/subscriptions/54707983-993e-43de-8d94-074451394eda/resourceGroups/adlsrg/providers/Microsoft.Storage/storageAccounts/adlsstorageaccount"))
                .withPolicyInfo(new PolicyInfo().withPolicyId(
                    "/subscriptions/54707983-993e-43de-8d94-074451394eda/resourceGroups/adlsrg/providers/Microsoft.DataProtection/backupVaults/adlsvault/backupPolicies/adlspolicy")
                    .withPolicyParameters(new PolicyParameters().withBackupDatasourceParametersList(Arrays.asList(
                        new AdlsBlobBackupDatasourceParameters().withContainersList(Arrays.asList("container1"))))))
                .withObjectType("BackupInstance"))
            .create();
    }

    // Use "Map.of" if available
    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}
