/** Samples for Assets Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/assets-delete.json
     */
    /**
     * Sample code: Delete an Asset.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void deleteAnAsset(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager
            .assets()
            .deleteWithResponse("contoso", "contosomedia", "ClimbingMountAdams", com.azure.core.util.Context.NONE);
    }
}
