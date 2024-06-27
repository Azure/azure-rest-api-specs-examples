
import com.azure.resourcemanager.scvmm.fluent.models.GuestAgentInner;
import com.azure.resourcemanager.scvmm.models.GuestAgentProperties;
import com.azure.resourcemanager.scvmm.models.GuestCredential;
import com.azure.resourcemanager.scvmm.models.HttpProxyConfiguration;
import com.azure.resourcemanager.scvmm.models.ProvisioningAction;

/**
 * Samples for GuestAgents Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/GuestAgents_Create_MaximumSet_Gen
     * .json
     */
    /**
     * Sample code: GuestAgents_Create_MaximumSet.
     * 
     * @param manager Entry point to ScvmmManager.
     */
    public static void guestAgentsCreateMaximumSet(com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        manager.guestAgents().create("gtgclehcbsyave", new GuestAgentInner().withProperties(new GuestAgentProperties()
            .withCredentials(new GuestCredential().withUsername("jqxuwirrcpfv").withPassword("fakeTokenPlaceholder"))
            .withHttpProxyConfig(new HttpProxyConfiguration().withHttpsProxy("uoyzyticmohohomlkwct"))
            .withProvisioningAction(ProvisioningAction.INSTALL)), com.azure.core.util.Context.NONE);
    }
}
