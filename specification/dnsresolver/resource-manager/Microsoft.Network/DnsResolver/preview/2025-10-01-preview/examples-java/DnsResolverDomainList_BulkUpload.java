
import com.azure.resourcemanager.dnsresolver.models.Action;
import com.azure.resourcemanager.dnsresolver.models.DnsResolverDomainListBulk;

/**
 * Samples for DnsResolverDomainLists Bulk.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/DnsResolverDomainList_BulkUpload.json
     */
    /**
     * Sample code: Upload DNS resolver domain list domains.
     * 
     * @param manager Entry point to DnsResolverManager.
     */
    public static void
        uploadDNSResolverDomainListDomains(com.azure.resourcemanager.dnsresolver.DnsResolverManager manager) {
        manager.dnsResolverDomainLists().bulk("sampleResourceGroup", "sampleDnsResolverDomainList",
            new DnsResolverDomainListBulk().withStorageUrl(
                "https://sampleStorageAccount.blob.core.windows.net/sample-container/sampleBlob.txt?sv=2022-11-02&sr=b&sig=39Up9jzHkxhUIhFEjEh9594DJxe7w6cIRCgOV6ICGS0%3A377&sp=rcw")
                .withAction(Action.UPLOAD),
            null, null, com.azure.core.util.Context.NONE);
    }
}
