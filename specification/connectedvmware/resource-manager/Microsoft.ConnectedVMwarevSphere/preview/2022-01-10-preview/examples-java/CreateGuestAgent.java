import com.azure.resourcemanager.connectedvmware.models.GuestCredential;
import com.azure.resourcemanager.connectedvmware.models.HttpProxyConfiguration;
import com.azure.resourcemanager.connectedvmware.models.ProvisioningAction;

/** Samples for GuestAgents Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2022-01-10-preview/examples/CreateGuestAgent.json
     */
    /**
     * Sample code: CreateGuestAgent.
     *
     * @param manager Entry point to ConnectedVMwareManager.
     */
    public static void createGuestAgent(com.azure.resourcemanager.connectedvmware.ConnectedVMwareManager manager) {
        manager
            .guestAgents()
            .define("default")
            .withExistingVirtualMachine("testrg", "ContosoVm")
            .withCredentials(new GuestCredential().withUsername("tempuser").withPassword("<password>"))
            .withHttpProxyConfig(new HttpProxyConfiguration().withHttpsProxy("http://192.1.2.3:8080"))
            .withProvisioningAction(ProvisioningAction.INSTALL)
            .create();
    }
}
