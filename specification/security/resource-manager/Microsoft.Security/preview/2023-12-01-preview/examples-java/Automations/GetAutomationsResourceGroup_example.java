
/**
 * Samples for Automations ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/preview/2023-12-01-preview/examples/Automations/
     * GetAutomationsResourceGroup_example.json
     */
    /**
     * Sample code: List all security automations of a specified resource group.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void listAllSecurityAutomationsOfASpecifiedResourceGroup(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.automations().listByResourceGroup("exampleResourceGroup", com.azure.core.util.Context.NONE);
    }
}
