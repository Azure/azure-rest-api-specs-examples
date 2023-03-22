/** Samples for AdaptiveApplicationControls Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/ApplicationWhitelistings/GetAdaptiveApplicationControlsGroup_example.json
     */
    /**
     * Sample code: Gets a configured application control VM/server group.
     *
     * @param manager Entry point to SecurityManager.
     */
    public static void getsAConfiguredApplicationControlVMServerGroup(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager
            .adaptiveApplicationControls()
            .getWithResponse("centralus", "ERELGROUP1", com.azure.core.util.Context.NONE);
    }
}
