Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-datafactory_1.0.0-beta.12/sdk/datafactory/azure-resourcemanager-datafactory/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.resourcemanager.datafactory.models.PrivateLinkConnectionApprovalRequest;
import com.azure.resourcemanager.datafactory.models.PrivateLinkConnectionState;

/** Samples for PrivateEndpointConnectionOperation CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/ApproveRejectPrivateEndpointConnection.json
     */
    /**
     * Sample code: Approves or rejects a private endpoint connection for a factory.
     *
     * @param manager Entry point to DataFactoryManager.
     */
    public static void approvesOrRejectsAPrivateEndpointConnectionForAFactory(
        com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager
            .privateEndpointConnectionOperations()
            .define("connection")
            .withExistingFactory("exampleResourceGroup", "exampleFactoryName")
            .withProperties(
                new PrivateLinkConnectionApprovalRequest()
                    .withPrivateLinkServiceConnectionState(
                        new PrivateLinkConnectionState()
                            .withStatus("Approved")
                            .withDescription("Approved by admin.")
                            .withActionsRequired("")))
            .create();
    }
}
```
