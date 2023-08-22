/** Samples for Creators Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/maps/resource-manager/Microsoft.Maps/stable/2023-06-01/examples/GetMapsCreator.json
     */
    /**
     * Sample code: Get Creator Resource.
     *
     * @param manager Entry point to AzureMapsManager.
     */
    public static void getCreatorResource(com.azure.resourcemanager.maps.AzureMapsManager manager) {
        manager
            .creators()
            .getWithResponse("myResourceGroup", "myMapsAccount", "myCreator", com.azure.core.util.Context.NONE);
    }
}
