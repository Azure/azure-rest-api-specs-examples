/** Samples for NetworkExperimentProfiles Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/frontdoor/resource-manager/Microsoft.Network/stable/2019-11-01/examples/NetworkExperimentDeleteProfile.json
     */
    /**
     * Sample code: Deletes an NetworkExperiment Profile by ProfileName.
     *
     * @param manager Entry point to FrontDoorManager.
     */
    public static void deletesAnNetworkExperimentProfileByProfileName(
        com.azure.resourcemanager.frontdoor.FrontDoorManager manager) {
        manager.networkExperimentProfiles().delete("MyResourceGroup", "MyProfile", com.azure.core.util.Context.NONE);
    }
}
