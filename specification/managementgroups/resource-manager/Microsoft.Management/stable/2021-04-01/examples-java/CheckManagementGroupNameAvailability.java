import com.azure.resourcemanager.managementgroups.models.CheckNameAvailabilityRequest;
import com.azure.resourcemanager.managementgroups.models.Type;

/** Samples for ResourceProvider CheckNameAvailability. */
public final class Main {
    /*
     * x-ms-original-file: specification/managementgroups/resource-manager/Microsoft.Management/stable/2021-04-01/examples/CheckManagementGroupNameAvailability.json
     */
    /**
     * Sample code: CheckManagementGroupNameAvailability.
     *
     * @param manager Entry point to ManagementGroupsManager.
     */
    public static void checkManagementGroupNameAvailability(
        com.azure.resourcemanager.managementgroups.ManagementGroupsManager manager) {
        manager
            .resourceProviders()
            .checkNameAvailabilityWithResponse(
                new CheckNameAvailabilityRequest()
                    .withName("nameTocheck")
                    .withType(Type.MICROSOFT_MANAGEMENT_MANAGEMENT_GROUPS),
                com.azure.core.util.Context.NONE);
    }
}
