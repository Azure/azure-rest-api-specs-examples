import com.azure.core.util.Context;
import com.azure.resourcemanager.labservices.models.EnableState;
import com.azure.resourcemanager.labservices.models.Image;

/** Samples for Images Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/labservices/resource-manager/Microsoft.LabServices/preview/2021-11-15-preview/examples/Images/patchImage.json
     */
    /**
     * Sample code: patchImage.
     *
     * @param manager Entry point to LabServicesManager.
     */
    public static void patchImage(com.azure.resourcemanager.labservices.LabServicesManager manager) {
        Image resource =
            manager.images().getWithResponse("testrg123", "testlabplan", "image1", Context.NONE).getValue();
        resource.update().withEnabledState(EnableState.ENABLED).apply();
    }
}
