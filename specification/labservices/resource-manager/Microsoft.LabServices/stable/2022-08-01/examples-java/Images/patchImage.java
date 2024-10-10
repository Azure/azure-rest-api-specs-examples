
import com.azure.resourcemanager.labservices.models.EnableState;
import com.azure.resourcemanager.labservices.models.Image;

/**
 * Samples for Images Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/labservices/resource-manager/Microsoft.LabServices/stable/2022-08-01/examples/Images/patchImage.
     * json
     */
    /**
     * Sample code: patchImage.
     * 
     * @param manager Entry point to LabServicesManager.
     */
    public static void patchImage(com.azure.resourcemanager.labservices.LabServicesManager manager) {
        Image resource = manager.images()
            .getWithResponse("testrg123", "testlabplan", "image1", com.azure.core.util.Context.NONE).getValue();
        resource.update().withEnabledState(EnableState.ENABLED).apply();
    }
}
