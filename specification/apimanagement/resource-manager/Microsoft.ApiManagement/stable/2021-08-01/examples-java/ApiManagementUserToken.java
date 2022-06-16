import com.azure.core.util.Context;
import com.azure.resourcemanager.apimanagement.models.KeyType;
import com.azure.resourcemanager.apimanagement.models.UserTokenParameters;
import java.time.OffsetDateTime;

/** Samples for User GetSharedAccessToken. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementUserToken.json
     */
    /**
     * Sample code: ApiManagementUserToken.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementUserToken(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .users()
            .getSharedAccessTokenWithResponse(
                "rg1",
                "apimService1",
                "userId1718",
                new UserTokenParameters()
                    .withKeyType(KeyType.PRIMARY)
                    .withExpiry(OffsetDateTime.parse("2019-04-21T00:44:24.2845269Z")),
                Context.NONE);
    }
}
