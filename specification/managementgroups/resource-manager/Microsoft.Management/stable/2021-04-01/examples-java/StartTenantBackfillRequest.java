/** Samples for ResourceProvider StartTenantBackfill. */
public final class Main {
    /*
     * x-ms-original-file: specification/managementgroups/resource-manager/Microsoft.Management/stable/2021-04-01/examples/StartTenantBackfillRequest.json
     */
    /**
     * Sample code: StartTenantBackfill.
     *
     * @param manager Entry point to ManagementGroupsManager.
     */
    public static void startTenantBackfill(com.azure.resourcemanager.managementgroups.ManagementGroupsManager manager) {
        manager.resourceProviders().startTenantBackfillWithResponse(com.azure.core.util.Context.NONE);
    }
}
