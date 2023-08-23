/** Samples for Creators Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/maps/resource-manager/Microsoft.Maps/stable/2023-06-01/examples/DeleteMapsCreator.json
     */
    /**
     * Sample code: Delete Creator Resource.
     *
     * @param manager Entry point to AzureMapsManager.
     */
    public static void deleteCreatorResource(com.azure.resourcemanager.maps.AzureMapsManager manager) {
        manager
            .creators()
            .deleteWithResponse("myResourceGroup", "myMapsAccount", "myCreator", com.azure.core.util.Context.NONE);
    }
}
