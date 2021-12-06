Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-recoveryservicesbackup_1.0.0-beta.2/sdk/recoveryservicesbackup/azure-resourcemanager-recoveryservicesbackup/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for PrivateEndpoint GetOperationStatus. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-07-01/examples/PrivateEndpointConnection/GetPrivateEndpointConnectionOperationStatus.json
     */
    /**
     * Sample code: Get OperationStatus.
     *
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void getOperationStatus(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager
            .privateEndpoints()
            .getOperationStatusWithResponse(
                "gaallavaultbvtd2msi",
                "gaallaRG",
                "gaallatestpe2.5704c932-249a-490b-a142-1396838cd3b",
                "0f48183b-0a44-4dca-aec1-bba5daab888a",
                Context.NONE);
    }
}
```
