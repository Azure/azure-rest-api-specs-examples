import com.azure.core.util.Context;
import com.azure.resourcemanager.synapse.models.CheckNameAvailabilityRequest;

/** Samples for Operations CheckNameAvailability. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/CheckNameAvailabilityWorkspaceAvailable.json
     */
    /**
     * Sample code: Check for a workspace name that is available.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void checkForAWorkspaceNameThatIsAvailable(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .operations()
            .checkNameAvailabilityWithResponse(
                new CheckNameAvailabilityRequest()
                    .withName("workspace1")
                    .withType("Microsoft.ProjectArcadia/workspaces"),
                Context.NONE);
    }
}
