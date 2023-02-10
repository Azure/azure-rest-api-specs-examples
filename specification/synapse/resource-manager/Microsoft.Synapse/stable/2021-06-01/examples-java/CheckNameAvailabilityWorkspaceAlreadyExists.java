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
                com.azure.core.util.Context.NONE);
    }
}
