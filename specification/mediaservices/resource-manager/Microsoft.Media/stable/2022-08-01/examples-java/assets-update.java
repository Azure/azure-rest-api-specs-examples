import com.azure.core.util.Context;
import com.azure.resourcemanager.mediaservices.models.Asset;

/** Samples for Assets Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2022-08-01/examples/assets-update.json
     */
    /**
     * Sample code: Update an Asset.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void updateAnAsset(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        Asset resource =
            manager.assets().getWithResponse("contoso", "contosomedia", "ClimbingMountBaker", Context.NONE).getValue();
        resource.update().withDescription("A documentary showing the ascent of Mount Baker in HD").apply();
    }
}
