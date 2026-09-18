
import com.azure.resourcemanager.dataprotection.models.AzureOperationalStoreParameters;
import com.azure.resourcemanager.dataprotection.models.BackupInstance;
import com.azure.resourcemanager.dataprotection.models.DataStoreTypes;
import com.azure.resourcemanager.dataprotection.models.Datasource;
import com.azure.resourcemanager.dataprotection.models.DatasourceSet;
import com.azure.resourcemanager.dataprotection.models.GenericBackupDatasourceParameters;
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
     * x-ms-original-file: 2026-06-01/BackupInstanceOperations/PutBackupInstanceWithGenericParameters.json
     */
    /**
     * Sample code: Create BackupInstance with GenericBackupDatasourceParameters.
     * 
     * @param manager Entry point to DataProtectionManager.
     */
    public static void createBackupInstanceWithGenericBackupDatasourceParameters(
        com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager.backupInstances().define("esan-volgroup-bi").withExistingBackupVault("ESAN-ECYBVTRG", "ESANVault")
            .withTags(mapOf("key1", "fakeTokenPlaceholder"))
            .withProperties(new BackupInstance().withFriendlyName("esan-volgroup-bi")
                .withDataSourceInfo(new Datasource().withDatasourceType("Microsoft.ElasticSan/elasticSans/volumeGroups")
                    .withObjectType("Datasource")
                    .withResourceId(
                        "/subscriptions/97cda027-4279-4cde-b4ff-19afa0021d87/resourceGroups/ESAN-ECYBVTRG/providers/Microsoft.ElasticSan/elasticSans/ecy-bvt-adhoc/volumeGroups/esan-volgroup")
                    .withResourceLocation("eastus2euap").withResourceName("esan-volgroup-bi")
                    .withResourceType("Microsoft.ElasticSan/elasticSans/volumeGroups")
                    .withResourceUri("SampleresourceUri123"))
                .withDataSourceSetInfo(new DatasourceSet()
                    .withDatasourceType("Microsoft.ElasticSan/elasticSans/volumeGroups").withObjectType("DatasourceSet")
                    .withResourceId(
                        "/subscriptions/97cda027-4279-4cde-b4ff-19afa0021d87/resourceGroups/ESAN-ECYBVTRG/providers/Microsoft.ElasticSan/elasticSans/ecy-bvt-adhoc")
                    .withResourceLocation("eastus2euap").withResourceType("Microsoft.ElasticSan/elasticSans"))
                .withPolicyInfo(new PolicyInfo().withPolicyId(
                    "/subscriptions/97cda027-4279-4cde-b4ff-19afa0021d87/resourceGroups/ESAN-ECYBVTRG/providers/Microsoft.DataProtection/backupVaults/ESANVault/backupPolicies/BVTPolicy")
                    .withPolicyParameters(new PolicyParameters()
                        .withDataStoreParametersList(Arrays.asList(new AzureOperationalStoreParameters()
                            .withDataStoreType(DataStoreTypes.OPERATIONAL_STORE).withResourceGroupId(
                                "/subscriptions/97cda027-4279-4cde-b4ff-19afa0021d87/resourceGroups/ESAN-ECYBVTRG")))
                        .withBackupDatasourceParametersList(Arrays.asList(new GenericBackupDatasourceParameters()
                            .withResourceSelectors(Arrays.asList("vol1", "vol2", "vol3"))))))
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
