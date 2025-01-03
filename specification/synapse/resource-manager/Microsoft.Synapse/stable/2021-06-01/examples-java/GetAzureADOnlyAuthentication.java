
import com.azure.resourcemanager.synapse.models.AzureADOnlyAuthenticationName;

/**
 * Samples for AzureADOnlyAuthentications Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/GetAzureADOnlyAuthentication.
     * json
     */
    /**
     * Sample code: Get Azure Active Directory Only Authentication property.
     * 
     * @param manager Entry point to SynapseManager.
     */
    public static void
        getAzureActiveDirectoryOnlyAuthenticationProperty(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager.azureADOnlyAuthentications().getWithResponse("workspace-6852", "workspace-2080",
            AzureADOnlyAuthenticationName.DEFAULT, com.azure.core.util.Context.NONE);
    }
}
