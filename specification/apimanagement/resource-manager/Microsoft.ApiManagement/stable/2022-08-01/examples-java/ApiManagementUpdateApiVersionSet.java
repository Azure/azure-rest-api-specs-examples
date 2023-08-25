import com.azure.resourcemanager.apimanagement.models.ApiVersionSetContract;
import com.azure.resourcemanager.apimanagement.models.VersioningScheme;

/** Samples for ApiVersionSet Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementUpdateApiVersionSet.json
     */
    /**
     * Sample code: ApiManagementUpdateApiVersionSet.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementUpdateApiVersionSet(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        ApiVersionSetContract resource =
            manager
                .apiVersionSets()
                .getWithResponse("rg1", "apimService1", "vs1", com.azure.core.util.Context.NONE)
                .getValue();
        resource
            .update()
            .withDisplayName("api set 1")
            .withVersioningScheme(VersioningScheme.SEGMENT)
            .withDescription("Version configuration")
            .withIfMatch("*")
            .apply();
    }
}
