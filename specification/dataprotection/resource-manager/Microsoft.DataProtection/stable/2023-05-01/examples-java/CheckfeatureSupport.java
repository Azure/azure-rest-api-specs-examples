import com.azure.resourcemanager.dataprotection.models.FeatureType;
import com.azure.resourcemanager.dataprotection.models.FeatureValidationRequest;

/** Samples for DataProtection CheckFeatureSupport. */
public final class Main {
    /*
     * x-ms-original-file: specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2023-05-01/examples/CheckfeatureSupport.json
     */
    /**
     * Sample code: Check Azure Vm Backup Feature Support.
     *
     * @param manager Entry point to DataProtectionManager.
     */
    public static void checkAzureVmBackupFeatureSupport(
        com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager
            .dataProtections()
            .checkFeatureSupportWithResponse(
                "WestUS",
                new FeatureValidationRequest().withFeatureType(FeatureType.DATA_SOURCE_TYPE),
                com.azure.core.util.Context.NONE);
    }
}
