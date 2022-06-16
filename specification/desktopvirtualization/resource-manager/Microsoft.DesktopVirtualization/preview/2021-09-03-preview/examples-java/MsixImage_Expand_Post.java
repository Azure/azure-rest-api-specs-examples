import com.azure.core.util.Context;
import com.azure.resourcemanager.desktopvirtualization.models.MsixImageUri;

/** Samples for MsixImages Expand. */
public final class Main {
    /*
     * x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/preview/2021-09-03-preview/examples/MsixImage_Expand_Post.json
     */
    /**
     * Sample code: MsixImage_Expand.
     *
     * @param manager Entry point to DesktopVirtualizationManager.
     */
    public static void msixImageExpand(
        com.azure.resourcemanager.desktopvirtualization.DesktopVirtualizationManager manager) {
        manager
            .msixImages()
            .expand("resourceGroup1", "hostpool1", new MsixImageUri().withUri("imagepath"), Context.NONE);
    }
}
