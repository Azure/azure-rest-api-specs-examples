/** Samples for ApiCollectionOffboarding Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2022-11-20-preview/examples/ApiCollections/APICollectionOffboarding_Delete_example.json
     */
    /**
     * Sample code: Delete a security recommendation task on a resource.
     *
     * @param manager Entry point to SecurityManager.
     */
    public static void deleteASecurityRecommendationTaskOnAResource(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager
            .apiCollectionOffboardings()
            .deleteWithResponse("rg1", "apimService1", "echo-api", com.azure.core.util.Context.NONE);
    }
}
