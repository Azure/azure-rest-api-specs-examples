import com.azure.resourcemanager.apimanagement.models.BackendAuthorizationHeaderCredentials;
import com.azure.resourcemanager.apimanagement.models.BackendCredentialsContract;
import com.azure.resourcemanager.apimanagement.models.BackendProtocol;
import com.azure.resourcemanager.apimanagement.models.BackendProxyContract;
import com.azure.resourcemanager.apimanagement.models.BackendTlsProperties;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for Backend CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateBackendProxyBackend.json
     */
    /**
     * Sample code: ApiManagementCreateBackendProxyBackend.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateBackendProxyBackend(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .backends()
            .define("proxybackend")
            .withExistingService("rg1", "apimService1")
            .withUrl("https://backendname2644/")
            .withProtocol(BackendProtocol.HTTP)
            .withDescription("description5308")
            .withCredentials(
                new BackendCredentialsContract()
                    .withQuery(mapOf("sv", Arrays.asList("xx", "bb", "cc")))
                    .withHeaderProperty(mapOf("x-my-1", Arrays.asList("val1", "val2")))
                    .withAuthorization(
                        new BackendAuthorizationHeaderCredentials().withScheme("Basic").withParameter("opensesma")))
            .withProxy(
                new BackendProxyContract()
                    .withUrl("http://192.168.1.1:8080")
                    .withUsername("Contoso\\admin")
                    .withPassword("fakeTokenPlaceholder"))
            .withTls(new BackendTlsProperties().withValidateCertificateChain(true).withValidateCertificateName(true))
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
