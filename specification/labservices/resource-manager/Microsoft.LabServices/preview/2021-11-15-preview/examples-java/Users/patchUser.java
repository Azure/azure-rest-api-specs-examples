import com.azure.core.util.Context;
import com.azure.resourcemanager.labservices.models.User;
import java.time.Duration;

/** Samples for Users Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/labservices/resource-manager/Microsoft.LabServices/preview/2021-11-15-preview/examples/Users/patchUser.json
     */
    /**
     * Sample code: patchUser.
     *
     * @param manager Entry point to LabServicesManager.
     */
    public static void patchUser(com.azure.resourcemanager.labservices.LabServicesManager manager) {
        User resource = manager.users().getWithResponse("testrg123", "testlab", "testuser", Context.NONE).getValue();
        resource.update().withAdditionalUsageQuota(Duration.parse("20:00")).apply();
    }
}
