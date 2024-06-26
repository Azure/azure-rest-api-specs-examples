import com.azure.core.util.Context;

/** Samples for TopLevelDomains List. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.DomainRegistration/stable/2022-03-01/examples/ListTopLevelDomains.json
     */
    /**
     * Sample code: List Top Level Domains.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listTopLevelDomains(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getTopLevelDomains().list(Context.NONE);
    }
}
