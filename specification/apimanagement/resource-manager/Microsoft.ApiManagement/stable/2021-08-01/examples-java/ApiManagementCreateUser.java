import com.azure.resourcemanager.apimanagement.models.Confirmation;

/** Samples for User CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementCreateUser.json
     */
    /**
     * Sample code: ApiManagementCreateUser.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateUser(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .users()
            .define("5931a75ae4bbd512288c680b")
            .withExistingService("rg1", "apimService1")
            .withEmail("foobar@outlook.com")
            .withFirstName("foo")
            .withLastName("bar")
            .withConfirmation(Confirmation.SIGNUP)
            .create();
    }
}
