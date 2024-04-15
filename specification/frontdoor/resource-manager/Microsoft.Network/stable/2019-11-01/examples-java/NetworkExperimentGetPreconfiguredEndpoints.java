
/**
 * Samples for PreconfiguredEndpoints List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/frontdoor/resource-manager/Microsoft.Network/stable/2019-11-01/examples/
     * NetworkExperimentGetPreconfiguredEndpoints.json
     */
    /**
     * Sample code: Gets a list of Preconfigured Endpoints.
     * 
     * @param manager Entry point to FrontDoorManager.
     */
    public static void getsAListOfPreconfiguredEndpoints(com.azure.resourcemanager.frontdoor.FrontDoorManager manager) {
        manager.preconfiguredEndpoints().list("MyResourceGroup", "MyProfile", com.azure.core.util.Context.NONE);
    }
}
