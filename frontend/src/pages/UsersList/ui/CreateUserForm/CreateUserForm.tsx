import {observer} from "mobx-react-lite";
import * as yup from "yup";
import {useMemo} from "react";
import {useFormik} from "formik";
import {Button, Grid, TextField} from "@mui/material";

import {
    emailValidation,
    commonStringValidation,
} from "shared/validations";
import {TUserWithoutID} from "shared/types";

import {UsersModel} from "../../model";
import styles from "./Styles.module.scss";
import {useTranslation} from "react-i18next";

function CreateUserForm({hideModal}: any) {
    const { t } = useTranslation();

    const validationSchema = useMemo(
        () =>
            yup.object().shape({
                name: commonStringValidation("Name", 3),
                surname: commonStringValidation("Surname", 3),
                email: emailValidation(),
            }),
        []
    );

    const {handleSubmit, values, handleChange, touched, errors} =
        useFormik<TUserWithoutID>({
            initialValues: {
                name: "",
                surname: "",
                email: "",
            },
            validationSchema,
            onSubmit: (value: TUserWithoutID) => {
                UsersModel.create(value);
                hideModal();
            },
        });

    return (
        <form onSubmit={handleSubmit} className={styles.centered}>
            <Grid item container spacing={2} direction="column">
                <Grid item>
                    <TextField
                        fullWidth
                        id="name"
                        name="name"
                        label={t("user:name")}
                        value={values.name}
                        onChange={handleChange}
                        error={touched.name && Boolean(errors.name)}
                        helperText={touched.name && errors.name}
                    />
                </Grid>
                <Grid item>
                    <TextField
                        fullWidth
                        id="surname"
                        name="surname"
                        label={"Surname"}
                        value={values.surname}
                        onChange={handleChange}
                        error={touched.surname && Boolean(errors.surname)}
                        helperText={touched.surname && errors.surname}
                    />
                </Grid>
                <Grid item>
                    <TextField
                        fullWidth
                        id="email"
                        name="email"
                        label={"Email"}
                        value={values.email}
                        onChange={handleChange}
                        error={touched.email && Boolean(errors.email)}
                        helperText={touched.email && errors.email}
                    />
                </Grid>
                <Grid item>
                    <Button color="primary" variant="contained" fullWidth type="submit">
                        {"Confirm"}
                    </Button>
                </Grid>
            </Grid>
        </form>
    );
}

export default observer(CreateUserForm);
