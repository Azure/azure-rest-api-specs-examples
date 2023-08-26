import com.azure.resourcemanager.apimanagement.models.ApiType;
import com.azure.resourcemanager.apimanagement.models.Protocol;
import java.util.Arrays;

/** Samples for Api CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateGraphQLApi.json
     */
    /**
     * Sample code: ApiManagementCreateGraphQLApi.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateGraphQLApi(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .apis()
            .define("tempgroup")
            .withExistingService("rg1", "apimService1")
            .withDisplayName("apiname1463")
            .withServiceUrl("https://api.spacex.land/graphql")
            .withPath("graphql-api")
            .withProtocols(Arrays.asList(Protocol.HTTP, Protocol.HTTPS))
            .withDescription("apidescription5200")
            .withApiType(ApiType.GRAPHQL)
            .create();
    }
}
