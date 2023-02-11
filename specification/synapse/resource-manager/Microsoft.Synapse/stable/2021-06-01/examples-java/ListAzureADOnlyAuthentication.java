/** Samples for AzureADOnlyAuthentications List. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/ListAzureADOnlyAuthentication.json
     */
    /**
     * Sample code: Get a list of Azure Active Directory Only Authentication property.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void getAListOfAzureActiveDirectoryOnlyAuthenticationProperty(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager.azureADOnlyAuthentications().list("workspace-6852", "workspace-2080", com.azure.core.util.Context.NONE);
    }
}
