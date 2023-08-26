import com.azure.resourcemanager.apimanagement.models.UserContract;

/** Samples for User Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementUpdateUser.json
     */
    /**
     * Sample code: ApiManagementUpdateUser.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementUpdateUser(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        UserContract resource =
            manager
                .users()
                .getWithResponse("rg1", "apimService1", "5931a75ae4bbd512a88c680b", com.azure.core.util.Context.NONE)
                .getValue();
        resource
            .update()
            .withEmail("foobar@outlook.com")
            .withFirstName("foo")
            .withLastName("bar")
            .withIfMatch("*")
            .apply();
    }
}
