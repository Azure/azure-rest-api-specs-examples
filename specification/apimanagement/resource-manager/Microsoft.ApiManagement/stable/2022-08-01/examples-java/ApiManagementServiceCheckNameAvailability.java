import com.azure.resourcemanager.apimanagement.models.ApiManagementServiceCheckNameAvailabilityParameters;

/** Samples for ApiManagementService CheckNameAvailability. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementServiceCheckNameAvailability.json
     */
    /**
     * Sample code: ApiManagementServiceCheckNameAvailability.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementServiceCheckNameAvailability(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .apiManagementServices()
            .checkNameAvailabilityWithResponse(
                new ApiManagementServiceCheckNameAvailabilityParameters().withName("apimService1"),
                com.azure.core.util.Context.NONE);
    }
}
