/** Samples for SecuritySolutionsReferenceData ListByHomeRegion. */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/SecuritySolutionsReferenceData/GetSecuritySolutionsReferenceDataSubscriptionLocation_example.json
     */
    /**
     * Sample code: Get security solutions from a security data location.
     *
     * @param manager Entry point to SecurityManager.
     */
    public static void getSecuritySolutionsFromASecurityDataLocation(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager
            .securitySolutionsReferenceDatas()
            .listByHomeRegionWithResponse("westcentralus", com.azure.core.util.Context.NONE);
    }
}
