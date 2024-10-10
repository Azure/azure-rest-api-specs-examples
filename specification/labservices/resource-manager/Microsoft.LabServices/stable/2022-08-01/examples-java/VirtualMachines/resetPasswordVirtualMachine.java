
import com.azure.resourcemanager.labservices.models.ResetPasswordBody;

/**
 * Samples for VirtualMachines ResetPassword.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/labservices/resource-manager/Microsoft.LabServices/stable/2022-08-01/examples/VirtualMachines/
     * resetPasswordVirtualMachine.json
     */
    /**
     * Sample code: resetPasswordVirtualMachine.
     * 
     * @param manager Entry point to LabServicesManager.
     */
    public static void resetPasswordVirtualMachine(com.azure.resourcemanager.labservices.LabServicesManager manager) {
        manager.virtualMachines().resetPassword("testrg123", "testlab", "template",
            new ResetPasswordBody().withUsername("example-username").withPassword("fakeTokenPlaceholder"),
            com.azure.core.util.Context.NONE);
    }
}
