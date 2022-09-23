import com.azure.core.util.Context;

/** Samples for AdaptiveApplicationControls Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/ApplicationWhitelistings/DeleteAdaptiveApplicationControls_example.json
     */
    /**
     * Sample code: Delete an application control machine group.
     *
     * @param manager Entry point to SecurityManager.
     */
    public static void deleteAnApplicationControlMachineGroup(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.adaptiveApplicationControls().deleteWithResponse("centralus", "GROUP1", Context.NONE);
    }
}
