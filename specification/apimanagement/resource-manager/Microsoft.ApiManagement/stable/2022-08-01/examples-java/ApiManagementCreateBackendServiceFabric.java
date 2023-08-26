import com.azure.resourcemanager.apimanagement.models.BackendProperties;
import com.azure.resourcemanager.apimanagement.models.BackendProtocol;
import com.azure.resourcemanager.apimanagement.models.BackendServiceFabricClusterProperties;
import com.azure.resourcemanager.apimanagement.models.X509CertificateName;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for Backend CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateBackendServiceFabric.json
     */
    /**
     * Sample code: ApiManagementCreateBackendServiceFabric.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateBackendServiceFabric(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .backends()
            .define("sfbackend")
            .withExistingService("rg1", "apimService1")
            .withUrl("fabric:/mytestapp/mytestservice")
            .withProtocol(BackendProtocol.HTTP)
            .withDescription("Service Fabric Test App 1")
            .withProperties(
                new BackendProperties()
                    .withServiceFabricCluster(
                        new BackendServiceFabricClusterProperties()
                            .withClientCertificateId(
                                "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/certificates/cert1")
                            .withMaxPartitionResolutionRetries(5)
                            .withManagementEndpoints(Arrays.asList("https://somecluster.com"))
                            .withServerX509Names(
                                Arrays
                                    .asList(
                                        new X509CertificateName()
                                            .withName("ServerCommonName1")
                                            .withIssuerCertificateThumbprint("IssuerCertificateThumbprint1")))))
            .create();
    }

    // Use "Map.of" if available
    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}
