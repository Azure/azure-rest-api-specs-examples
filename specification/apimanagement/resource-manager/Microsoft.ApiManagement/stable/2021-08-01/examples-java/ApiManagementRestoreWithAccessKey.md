```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.apimanagement.models.AccessType;
import com.azure.resourcemanager.apimanagement.models.ApiManagementServiceBackupRestoreParameters;

/** Samples for ApiManagementService Restore. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementRestoreWithAccessKey.json
     */
    /**
     * Sample code: ApiManagementRestoreService.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementRestoreService(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .apiManagementServices()
            .restore(
                "rg1",
                "apimService1",
                new ApiManagementServiceBackupRestoreParameters()
                    .withStorageAccount("teststorageaccount")
                    .withContainerName("backupContainer")
                    .withBackupName("apimService1backup_2017_03_19")
                    .withAccessType(AccessType.ACCESS_KEY)
                    .withAccessKey("**************************************************"),
                Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-apimanagement_1.0.0-beta.3/sdk/apimanagement/azure-resourcemanager-apimanagement/README.md) on how to add the SDK to your project and authenticate.
