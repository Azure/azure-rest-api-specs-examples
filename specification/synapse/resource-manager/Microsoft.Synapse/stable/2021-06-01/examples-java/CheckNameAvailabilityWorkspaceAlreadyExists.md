Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-synapse_1.0.0-beta.4/sdk/synapse/azure-resourcemanager-synapse/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.synapse.models.CheckNameAvailabilityRequest;

/** Samples for Operations CheckNameAvailability. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/CheckNameAvailabilityWorkspaceAlreadyExists.json
     */
    /**
     * Sample code: Check for a workspace name that already exists.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void checkForAWorkspaceNameThatAlreadyExists(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .operations()
            .checkNameAvailabilityWithResponse(
                new CheckNameAvailabilityRequest().withName("workspace1").withType("Microsoft.Synapse/workspaces"),
                Context.NONE);
    }
}
```
