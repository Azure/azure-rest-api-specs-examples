Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-apimanagement_1.0.0-beta.3/sdk/apimanagement/azure-resourcemanager-apimanagement/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.apimanagement.models.AccessType;
import com.azure.resourcemanager.apimanagement.models.ApiManagementServiceBackupRestoreParameters;

/** Samples for ApiManagementService Backup. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementBackupWithUserAssignedManagedIdentity.json
     */
    /**
     * Sample code: ApiManagementBackupWithUserAssignedManagedIdentity.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementBackupWithUserAssignedManagedIdentity(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .apiManagementServices()
            .backup(
                "rg1",
                "apimService1",
                new ApiManagementServiceBackupRestoreParameters()
                    .withStorageAccount("contosorpstorage")
                    .withContainerName("apim-backups")
                    .withBackupName("backup5")
                    .withAccessType(AccessType.USER_ASSIGNED_MANAGED_IDENTITY)
                    .withClientId("XXXXX-a154-4830-XXXX-46a12da1a1e2"),
                Context.NONE);
    }
}
```
