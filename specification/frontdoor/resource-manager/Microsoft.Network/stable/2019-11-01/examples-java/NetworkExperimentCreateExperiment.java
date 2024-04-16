
import com.azure.resourcemanager.frontdoor.models.Endpoint;
import com.azure.resourcemanager.frontdoor.models.State;

/**
 * Samples for Experiments CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/frontdoor/resource-manager/Microsoft.Network/stable/2019-11-01/examples/
     * NetworkExperimentCreateExperiment.json
     */
    /**
     * Sample code: Creates an Experiment.
     * 
     * @param manager Entry point to FrontDoorManager.
     */
    public static void createsAnExperiment(com.azure.resourcemanager.frontdoor.FrontDoorManager manager) {
        manager.experiments().define("MyExperiment").withRegion((String) null)
            .withExistingNetworkExperimentProfile("MyResourceGroup", "MyProfile")
            .withDescription("this is my first experiment!")
            .withEndpointA(new Endpoint().withName("endpoint A").withEndpoint("endpointA.net"))
            .withEndpointB(new Endpoint().withName("endpoint B").withEndpoint("endpointB.net"))
            .withEnabledState(State.ENABLED).create();
    }
}
