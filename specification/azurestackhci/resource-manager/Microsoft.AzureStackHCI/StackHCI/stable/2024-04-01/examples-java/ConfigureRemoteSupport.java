
import com.azure.resourcemanager.azurestackhci.models.RemoteSupportRequest;
import com.azure.resourcemanager.azurestackhci.models.RemoteSupportRequestProperties;
import com.azure.resourcemanager.azurestackhci.models.RemoteSupportType;
import java.time.OffsetDateTime;

/**
 * Samples for Clusters ConfigureRemoteSupport.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/StackHCI/stable/2024-04-01/examples/
     * ConfigureRemoteSupport.json
     */
    /**
     * Sample code: Configure Remote Support.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void configureRemoteSupport(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.clusters().configureRemoteSupport("test-rg", "mycluster",
            new RemoteSupportRequest().withProperties(new RemoteSupportRequestProperties()
                .withExpirationTimestamp(OffsetDateTime.parse("2020-01-01T17:18:19.1234567Z"))
                .withRemoteSupportType(RemoteSupportType.ENABLE)),
            com.azure.core.util.Context.NONE);
    }
}
