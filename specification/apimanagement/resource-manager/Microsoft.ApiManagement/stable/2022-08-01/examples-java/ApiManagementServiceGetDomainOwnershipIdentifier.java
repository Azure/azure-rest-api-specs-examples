/** Samples for ApiManagementService GetDomainOwnershipIdentifier. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementServiceGetDomainOwnershipIdentifier.json
     */
    /**
     * Sample code: ApiManagementServiceGetDomainOwnershipIdentifier.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementServiceGetDomainOwnershipIdentifier(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.apiManagementServices().getDomainOwnershipIdentifierWithResponse(com.azure.core.util.Context.NONE);
    }
}
