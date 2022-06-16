/** Samples for Assets CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/assets-create.json
     */
    /**
     * Sample code: Create an Asset.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void createAnAsset(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager
            .assets()
            .define("ClimbingMountLogan")
            .withExistingMediaService("contoso", "contosomedia")
            .withDescription("A documentary showing the ascent of Mount Logan")
            .withStorageAccountName("storage0")
            .create();
    }
}
