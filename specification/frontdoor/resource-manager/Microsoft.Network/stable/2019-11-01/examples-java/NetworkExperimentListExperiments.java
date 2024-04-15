
/**
 * Samples for Experiments ListByProfile.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/frontdoor/resource-manager/Microsoft.Network/stable/2019-11-01/examples/
     * NetworkExperimentListExperiments.json
     */
    /**
     * Sample code: Gets a list of Experiments.
     * 
     * @param manager Entry point to FrontDoorManager.
     */
    public static void getsAListOfExperiments(com.azure.resourcemanager.frontdoor.FrontDoorManager manager) {
        manager.experiments().listByProfile("MyResourceGroup", "MyProfile", com.azure.core.util.Context.NONE);
    }
}
