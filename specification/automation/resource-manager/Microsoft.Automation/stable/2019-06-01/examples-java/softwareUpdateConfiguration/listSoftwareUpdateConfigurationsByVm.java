import com.azure.core.util.Context;

/** Samples for SoftwareUpdateConfigurations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2019-06-01/examples/softwareUpdateConfiguration/listSoftwareUpdateConfigurationsByVm.json
     */
    /**
     * Sample code: List software update configurations Targeting a specific azure virtual machine.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void listSoftwareUpdateConfigurationsTargetingASpecificAzureVirtualMachine(
        com.azure.resourcemanager.automation.AutomationManager manager) {
        manager
            .softwareUpdateConfigurations()
            .listWithResponse(
                "mygroup",
                "myaccount",
                null,
                "properties/updateConfiguration/azureVirtualMachines/any(m: m eq"
                    + " '/subscriptions/1a7d4044-286c-4acb-969a-96639265bf2e/resourceGroups/myresources/providers/Microsoft.Compute/virtualMachines/vm-01')",
                Context.NONE);
    }
}
