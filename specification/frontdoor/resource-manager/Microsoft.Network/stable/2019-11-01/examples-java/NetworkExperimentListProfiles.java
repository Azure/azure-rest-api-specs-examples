/** Samples for NetworkExperimentProfiles ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/frontdoor/resource-manager/Microsoft.Network/stable/2019-11-01/examples/NetworkExperimentListProfiles.json
     */
    /**
     * Sample code: List NetworkExperiment Profiles in a Resource Group.
     *
     * @param manager Entry point to FrontDoorManager.
     */
    public static void listNetworkExperimentProfilesInAResourceGroup(
        com.azure.resourcemanager.frontdoor.FrontDoorManager manager) {
        manager.networkExperimentProfiles().listByResourceGroup("MyResourceGroup", com.azure.core.util.Context.NONE);
    }
}
