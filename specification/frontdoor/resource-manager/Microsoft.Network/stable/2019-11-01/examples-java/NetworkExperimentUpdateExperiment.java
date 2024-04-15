
import com.azure.resourcemanager.frontdoor.models.Experiment;
import com.azure.resourcemanager.frontdoor.models.State;

/**
 * Samples for Experiments Update.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/frontdoor/resource-manager/Microsoft.Network/stable/2019-11-01/examples/
     * NetworkExperimentUpdateExperiment.json
     */
    /**
     * Sample code: Updates an Experiment.
     * 
     * @param manager Entry point to FrontDoorManager.
     */
    public static void updatesAnExperiment(com.azure.resourcemanager.frontdoor.FrontDoorManager manager) {
        Experiment resource = manager.experiments()
            .getWithResponse("MyResourceGroup", "MyProfile", "MyExperiment", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withDescription("string").withEnabledState(State.ENABLED).apply();
    }
}
