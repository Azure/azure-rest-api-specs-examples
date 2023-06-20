/** Samples for ResourceProvider TenantBackfillStatus. */
public final class Main {
    /*
     * x-ms-original-file: specification/managementgroups/resource-manager/Microsoft.Management/stable/2021-04-01/examples/TenantBackfillStatusRequest.json
     */
    /**
     * Sample code: TenantBackfillStatus.
     *
     * @param manager Entry point to ManagementGroupsManager.
     */
    public static void tenantBackfillStatus(
        com.azure.resourcemanager.managementgroups.ManagementGroupsManager manager) {
        manager.resourceProviders().tenantBackfillStatusWithResponse(com.azure.core.util.Context.NONE);
    }
}
