```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.dataprotection.models.FeatureType;
import com.azure.resourcemanager.dataprotection.models.FeatureValidationRequest;

/** Samples for DataProtection CheckFeatureSupport. */
public final class Main {
    /*
     * x-ms-original-file: specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2021-07-01/examples/CheckfeatureSupport.json
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
                "WestUS", new FeatureValidationRequest().withFeatureType(FeatureType.DATA_SOURCE_TYPE), Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-dataprotection_1.0.0-beta.1/sdk/dataprotection/azure-resourcemanager-dataprotection/README.md) on how to add the SDK to your project and authenticate.
