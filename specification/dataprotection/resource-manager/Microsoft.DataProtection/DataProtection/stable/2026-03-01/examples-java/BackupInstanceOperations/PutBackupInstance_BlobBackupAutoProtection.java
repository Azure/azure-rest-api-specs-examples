
import com.azure.resourcemanager.dataprotection.models.BackupInstance;
import com.azure.resourcemanager.dataprotection.models.BlobBackupAutoProtectionRule;
import com.azure.resourcemanager.dataprotection.models.BlobBackupDatasourceParametersForAutoProtection;
import com.azure.resourcemanager.dataprotection.models.BlobBackupPatternType;
import com.azure.resourcemanager.dataprotection.models.BlobBackupRuleBasedAutoProtectionSettings;
import com.azure.resourcemanager.dataprotection.models.BlobBackupRuleMode;
import com.azure.resourcemanager.dataprotection.models.Datasource;
import com.azure.resourcemanager.dataprotection.models.DatasourceSet;
import com.azure.resourcemanager.dataprotection.models.PolicyInfo;
import com.azure.resourcemanager.dataprotection.models.PolicyParameters;
import java.util.Arrays;

/**
 * Samples for BackupInstances CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/BackupInstanceOperations/PutBackupInstance_BlobBackupAutoProtection.json
     */
    /**
     * Sample code: Create BackupInstance With BlobBackupAutoProtection.
     * 
     * @param manager Entry point to DataProtectionManager.
     */
    public static void createBackupInstanceWithBlobBackupAutoProtection(
        com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager.backupInstances().define("blobstorageaccount-blobstorageaccount-2a76f8a-c176-4f7d-819e-95157e2b0071")
            .withExistingBackupVault("blobrg", "blobvault")
            .withProperties(new BackupInstance().withFriendlyName("blobstorageaccount\\blobbackupinstance")
                .withDataSourceInfo(new Datasource()
                    .withDatasourceType("Microsoft.Storage/storageAccounts/blobServices").withObjectType("Datasource")
                    .withResourceId(
                        "/subscriptions/54707983-993e-43de-8d94-074451394eda/resourceGroups/blobrg/providers/Microsoft.Storage/storageAccounts/blobstorageaccount")
                    .withResourceLocation("centraluseuap").withResourceName("blobstorageaccount")
                    .withResourceType("microsoft.storage/storageAccounts").withResourceUri(
                        "/subscriptions/54707983-993e-43de-8d94-074451394eda/resourceGroups/blobrg/providers/Microsoft.Storage/storageAccounts/blobstorageaccount"))
                .withDataSourceSetInfo(new DatasourceSet()
                    .withDatasourceType("Microsoft.Storage/storageAccounts/blobServices")
                    .withObjectType("DatasourceSet")
                    .withResourceId(
                        "/subscriptions/54707983-993e-43de-8d94-074451394eda/resourceGroups/blobrg/providers/Microsoft.Storage/storageAccounts/blobstorageaccount")
                    .withResourceLocation("centraluseuap").withResourceName("blobstorageaccount")
                    .withResourceType("microsoft.storage/storageAccounts").withResourceUri(
                        "/subscriptions/54707983-993e-43de-8d94-074451394eda/resourceGroups/blobrg/providers/Microsoft.Storage/storageAccounts/blobstorageaccount"))
                .withPolicyInfo(new PolicyInfo().withPolicyId(
                    "/subscriptions/54707983-993e-43de-8d94-074451394eda/resourceGroups/blobrg/providers/Microsoft.DataProtection/backupVaults/blobvault/backupPolicies/blobpolicy")
                    .withPolicyParameters(new PolicyParameters().withBackupDatasourceParametersList(
                        Arrays.asList(new BlobBackupDatasourceParametersForAutoProtection().withAutoProtectionSettings(
                            new BlobBackupRuleBasedAutoProtectionSettings().withEnabled(true)
                                .withRules(Arrays.asList(
                                    new BlobBackupAutoProtectionRule().withObjectType("BlobBackupAutoProtectionRule")
                                        .withMode(BlobBackupRuleMode.EXCLUDE).withType(BlobBackupPatternType.PREFIX)
                                        .withPattern("temp-"),
                                    new BlobBackupAutoProtectionRule().withObjectType("BlobBackupAutoProtectionRule")
                                        .withMode(BlobBackupRuleMode.EXCLUDE).withType(BlobBackupPatternType.PREFIX)
                                        .withPattern("test-"))))))))
                .withObjectType("BackupInstance"))
            .create();
    }
}
