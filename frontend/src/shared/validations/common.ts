import * as yup from 'yup'
import i18next from 'i18next'

export const commonStringValidation = (fieldName: string, minSymbols: number = 1) =>
    yup
        .string()
        .min(
            minSymbols,
            i18next.t('validation:error.minSymbols', {
                field: i18next.t(`user:${fieldName.toLowerCase()}`),
                count: minSymbols,
            })
        )
        .required(
            i18next.t('validation:error.isRequired', {
                field: i18next.t(`user:${fieldName.toLowerCase()}`),
            })
        )

export const emailValidation = () =>
    yup
        .string()
        .email(
            i18next.t('validation:error.validField', {
                field: i18next.t('user:email'),
            })
        )
        .required(
            i18next.t('validation:error.isRequired', {
                field: i18next.t('user:email'),
            })
        )
