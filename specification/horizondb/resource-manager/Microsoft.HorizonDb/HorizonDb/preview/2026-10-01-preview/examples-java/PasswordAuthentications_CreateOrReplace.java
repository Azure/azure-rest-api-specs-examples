
import com.azure.resourcemanager.horizondb.fluent.models.PasswordAuthenticationInner;
import com.azure.resourcemanager.horizondb.models.AuthenticationState;
import com.azure.resourcemanager.horizondb.models.PasswordAuthenticationProperties;

/**
 * Samples for PasswordAuthentications CreateOrReplace.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/PasswordAuthentications_CreateOrReplace.json
     */
    /**
     * Sample code: Enable password authentication and set the administrator login and password during initial setup.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void enablePasswordAuthenticationAndSetTheAdministratorLoginAndPasswordDuringInitialSetup(
        com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.passwordAuthentications().createOrReplace("example-resource-group", "example-cluster",
            new PasswordAuthenticationInner()
                .withProperties(new PasswordAuthenticationProperties().withState(AuthenticationState.ENABLED)
                    .withAdministratorLogin("adminlogin1").withAdministratorLoginPassword("fakeTokenPlaceholder")),
            com.azure.core.util.Context.NONE);
    }
}
