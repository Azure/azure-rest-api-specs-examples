```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.recoveryservicesbackup.models.PrivateEndpoint;
import com.azure.resourcemanager.recoveryservicesbackup.models.PrivateEndpointConnection;
import com.azure.resourcemanager.recoveryservicesbackup.models.PrivateEndpointConnectionResource;
import com.azure.resourcemanager.recoveryservicesbackup.models.PrivateEndpointConnectionStatus;
import com.azure.resourcemanager.recoveryservicesbackup.models.PrivateLinkServiceConnectionState;
import com.azure.resourcemanager.recoveryservicesbackup.models.ProvisioningState;

/** Samples for PrivateEndpointConnection Put. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-12-01/examples/PrivateEndpointConnection/PutPrivateEndpointConnection.json
     */
    /**
     * Sample code: Update PrivateEndpointConnection.
     *
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void updatePrivateEndpointConnection(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        PrivateEndpointConnectionResource resource =
            manager
                .privateEndpointConnections()
                .getWithResponse(
                    "gaallavaultbvtd2msi",
                    "gaallaRG",
                    "gaallatestpe2.5704c932-249a-490b-a142-1396838cd3b",
                    Context.NONE)
                .getValue();
        resource
            .update()
            .withProperties(
                new PrivateEndpointConnection()
                    .withProvisioningState(ProvisioningState.SUCCEEDED)
                    .withPrivateEndpoint(
                        new PrivateEndpoint()
                            .withId(
                                "/subscriptions/04cf684a-d41f-4550-9f70-7708a3a2283b/resourceGroups/gaallaRG/providers/Microsoft.Network/privateEndpoints/gaallatestpe3"))
                    .withPrivateLinkServiceConnectionState(
                        new PrivateLinkServiceConnectionState()
                            .withStatus(PrivateEndpointConnectionStatus.APPROVED)
                            .withDescription("Approved by johndoe@company.com")))
            .apply();
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-recoveryservicesbackup_1.0.0-beta.4/sdk/recoveryservicesbackup/azure-resourcemanager-recoveryservicesbackup/README.md) on how to add the SDK to your project and authenticate.
