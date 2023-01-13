/** Samples for NetworkExperimentProfiles GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/frontdoor/resource-manager/Microsoft.Network/stable/2019-11-01/examples/NetworkExperimentGetProfile.json
     */
    /**
     * Sample code: Gets an NetworkExperiment Profile by Profile Id.
     *
     * @param manager Entry point to FrontDoorManager.
     */
    public static void getsAnNetworkExperimentProfileByProfileId(
        com.azure.resourcemanager.frontdoor.FrontDoorManager manager) {
        manager
            .networkExperimentProfiles()
            .getByResourceGroupWithResponse("MyResourceGroup", "MyProfile", com.azure.core.util.Context.NONE);
    }
}
